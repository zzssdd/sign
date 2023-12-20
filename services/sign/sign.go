package sign

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"sign/dao/cache/model"
	model3 "sign/dao/db/model"
	model2 "sign/dao/mq/model"
	"sign/kitex_gen/sign/sign"
	. "sign/pkg/log"
	"sign/utils"
	"sync"
	"time"
)

// Sign implements the SignServiceImpl interface.
func (s *SignServiceImpl) Sign(ctx context.Context, req *sign.Empty) (resp *sign.Empty, err error) {
	resp = new(sign.Empty)
	msgChan := s.mq.ConsumeSignMsg()
	var signInfo *model2.Sign
	for msg := range msgChan {
		err = json.Unmarshal(msg.Body, &signInfo)
		if err != nil {
			logrus.Errorf("unmarshal sign msg error:%v\n", err)
			continue
		}
		err = s.HandleSign(signInfo)
		if err != nil {
			Log.Errorf("s.handleSign error:%v\n", err)
			_ = msg.Ack(false)
		} else {
			_ = msg.Ack(true)
		}
	}
	return
}

func (s *SignServiceImpl) HandleSign(signInfo *model2.Sign) error {
	wg := sync.WaitGroup{}
	var group *model3.Group
	var userGroups string
	wg.Add(2)
	go func() {
		if ok, err := s.cache.ExistAndExpireGroup(signInfo.Gid); err == nil && ok {
			getGroup, err := s.cache.GetGroup(signInfo.Gid)
			if err != nil {
				Log.Errorf("get group from cache error:%v\n", err)
				return
			}
			group = &model3.Group{
				Name:     getGroup.Name,
				Owner:    getGroup.Owner,
				Places:   getGroup.Places,
				Sign_in:  getGroup.SignIn,
				Sign_out: getGroup.SignOut,
				Count:    getGroup.Count,
				Score:    getGroup.Score,
			}
		} else {
			nums := rand.Int63()
			for !s.cache.AllLocker.SignLocker.Lock(nums) {
			}
			if ok, err := s.cache.ExistAndExpireGroup(signInfo.Gid); err == nil && ok {
				getGroup, err := s.cache.GetGroup(signInfo.Gid)
				if err != nil {
					Log.Errorf("get group from cache error:%v\n", err)
					return
				}
				group = &model3.Group{
					Name:     getGroup.Name,
					Owner:    getGroup.Owner,
					Places:   getGroup.Places,
					Sign_in:  getGroup.SignIn,
					Sign_out: getGroup.SignOut,
					Count:    getGroup.Count,
					Score:    getGroup.Score,
				}
			} else {
				group, err = s.db.Group.GetGroup(signInfo.Gid)
				if err != nil {
					Log.Errorf("get group from db error:%v\n", err)
					return
				}
				info := &model.Group{
					Name:    group.Name,
					Owner:   group.Owner,
					Places:  group.Places,
					SignIn:  group.Sign_in,
					SignOut: group.Sign_out,
					Count:   group.Count,
					Score:   group.Score,
				}
				_ = s.cache.StoreGroup(signInfo.Gid, info)
			}
			s.cache.SignLocker.UnLock(nums)
		}
		wg.Done()
	}()
	go func() {
		if is_ok, err := s.cache.ExistAndExpireUserGroups(signInfo.Uid); err == nil && is_ok {
			userGroups, err = s.cache.GetUserGroupsInfo(signInfo.Uid)
			if err != nil {
				Log.Errorf("get user group from cache error:%v\n", err)
				return
			}
		} else {
			userGroups, err = s.db.Group.GetUserGroups(signInfo.Uid)
			if err != nil {
				Log.Errorf("get user groups from db error:%v\n", err)
				return
			}
			err = s.cache.StoreUserGroupsInfo(signInfo.Uid, userGroups)
			if err != nil {
				Log.Errorf("store user group into cache error:%v\n", err)
			}
		}
		wg.Done()
	}()
	wg.Wait()
	if is_ok, err := s.cache.ExistSignPos(signInfo.Gid); err != nil || !is_ok {
		placesMap := utils.ParsePlacesString(group.Places)
		for _, v := range placesMap {
			info := &model.SignPos{
				Gid:        signInfo.Gid,
				Name:       v.Name,
				Latitle:    v.Longtitude,
				Longtitude: v.Latitude,
			}
			_ = s.cache.Sign.AddSignPos(info)
		}
	}
	signin, _ := time.Parse("15:04:05", group.Sign_in)
	signout, _ := time.Parse("15:04:05", group.Sign_out)
	signInTime, _ := time.Parse("15:04:05", signInfo.SignInTime)
	signOutTime, _ := time.Parse("15:04:05", signInfo.SignOutTime)
	if signInfo.Flag == 0 && signInTime.After(signin) {
		return fmt.Errorf("signin time too late")
	} else if signInfo.Flag == 1 && signOutTime.Before(signout) {
		return fmt.Errorf("signout time too early")
	}
	if !utils.StringContainInt64(userGroups, signInfo.Gid) {
		return fmt.Errorf("please join group first")
	}
	placeSlice := utils.ParsePlacesString(signInfo.Place)
	if len(placeSlice) <= 0 {
		return fmt.Errorf("sign pos not incorrect")
	}
	if is_ok, err := s.cache.JudgeSignPos(signInfo.Gid, placeSlice[0].Latitude, placeSlice[0].Longtitude); err != nil || !is_ok {
		return fmt.Errorf("not inside sign pos")
	}
	year, month, day := time.Now().Date()
	date := fmt.Sprintf("%d-%d-%d", year, month, day)
	if signInfo.Flag == 0 {
		info := &model3.SignDate{
			Uid:           signInfo.Uid,
			Gid:           signInfo.Gid,
			Date:          date,
			Signin_time:   signInfo.SignInTime,
			Signin_places: signInfo.Place,
		}
		err := s.db.Sign.StoreSignUserData(info)
		if err != nil {
			Log.Errorf("store sign date error:%v\n", err)
			return err
		}
		info2 := &model.Sign{
			Signin_time:   signInfo.SignInTime,
			Signin_places: signInfo.Place,
		}
		_ = s.cache.UserSign(signInfo.Uid, signInfo.Gid, date, info2)
	} else if signInfo.Flag == 1 {
		var userSign *model.Sign
		if is_ok, err := s.cache.ExistAndExpireSign(signInfo.Uid, signInfo.Gid, date); err == nil && is_ok {
			userSign, err = s.cache.GetUserSign(signInfo.Uid, signInfo.Gid, date)
			if err != nil {
				return err
			}
		} else {
			userDate, err := s.db.Sign.GetSignUserDate(signInfo.Uid, signInfo.Gid, date)
			if err != nil {
				return fmt.Errorf("get sign user Date error:%v\n", err)
			}
			userSign = &model.Sign{
				Signin_time:    userDate.Signin_time,
				Signin_places:  userDate.Signin_places,
				Signout_time:   userDate.Signout_time,
				Signout_places: userDate.Signout_places,
			}
			_ = s.cache.UserSign(signInfo.Uid, signInfo.Gid, date, userSign)
		}
		if userSign.Signin_time == "" {
			return fmt.Errorf("user not sign in")
		}
		if userSign.Signout_time == "" || userSign.Signout_time == "00:00:00" {
			err := s.db.User.AddUserScore(signInfo.Uid, group.Score)
			if err != nil {
				return err
			}
			_ = s.cache.DeleteUserScore(signInfo.Uid)
		}
		userSign.Signout_time = signInfo.SignOutTime
		userSign.Signout_places = signInfo.Place
		_ = s.cache.UserSign(signInfo.Uid, signInfo.Gid, date, userSign)
		info := &model3.SignDate{
			Uid:            signInfo.Uid,
			Gid:            signInfo.Gid,
			Date:           date,
			Signout_time:   signInfo.SignOutTime,
			Signout_places: signInfo.Place,
		}
		err := s.db.Sign.UpdateSignoutUserData(info)
		if err != nil {
			return err
		}
	}
	return nil
}
