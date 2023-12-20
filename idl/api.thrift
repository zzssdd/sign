include "base.thrift"

namespace go sign.api

struct JoinReq{
    1:i64 gid;
}

struct GroupInfo{
    1:string places;
    2:string sign_in;
    3:string sign_out;
    4:i32 score;
}

struct SignReq{
    1:i64 gid;
    2:optional string signinTime;
    3:optional string signoutTime;
    4:optional string place;
    5:i8 flag;
}

struct MonthSignReq{
    1:i64 gid;
    2:string month;
}

struct ChooseReq{
    1:i64 id;
    2:string chooseTime;
}

struct ChooseSubmitReq{
    1:i64 id;
}

service SignApi{
    base.BaseResp Register(1:base.UserInfo req)(api.post="/register")
    base.LoginResp Login(1:base.UserInfo req)(api.post="/login")
    base.BaseResp Join(1:JoinReq req)(api.post="/join")
    base.BaseResp CreateGroup(1:GroupInfo req)(api.post="/group")
    base.BaseResp Sign(1:SignReq req)(api.post="/sign")
    base.MonthSignResp SignMonth(1:MonthSignReq req)(api.get="/signMonth")
    base.ChooseResp Choose(1:ChooseReq req)(api.get="/choose")
    base.ChooseSubmitResp ChooseSubmit(1:ChooseSubmitReq req)(api.post="/choose")
    base.BaseResp AddPrize(1:base.PrizeInfo req)(api.post="/prize")
    base.BaseResp AddActivity(1:base.ActicityInfo req)(api.post="/activity")
}