package common

import "github.com/jinzhu/copier"

func BindRPC(req interface{}, rpcReq interface{}) error {
	if rpcReq != nil {
		err := copier.Copy(rpcReq, req)
		if err != nil {
			return err
		}
	}
	return nil
}
