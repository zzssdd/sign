namespace go sign.base

struct BaseResp{
    1:i32 code;
    2:string msg;
}

struct UserInfo{
    1:string email;
    2:string password;
}

struct LoginResp{
    1:BaseResp base;
    2:optional string token;
}

struct JoinReq{
    1:i64 uid;
    2:i64 gid;
}

struct GroupInfo{
    1:i64 owner;
    2:string places;
    3:string sign_in;
    4:string sign_out;
}

struct SignReq{
    1:i64 uid;
    2:i64 gid;
    3:optional string signinTime;
    4:optional string signoutTime;
    5:optional string Place;
}

struct MonthSignReq{
    1:i64 uid;
    2:i64 gid;
    3:string month;
}

struct MonthSignResp{
    1:BaseResp base;
    2:optional i32 data;
}

struct PrizeInfo{
    1:optional i32 id;
    2:string name;
    3:i64 gid;
}

struct ChooseReq{
    1:i64 uid;
    2:i64 id;
    3:string chooseTime;
}

struct ChooseResp{
    1:BaseResp base;
    2:optional i64 id;
}

struct ChooseSubmitReq{
    1:i64 uid;
    2:i64 id;
}

struct ChooseSubmitResp{
    1:BaseResp base;
    2:optional PrizeInfo info;
}

struct PrizePair{
    1:i32 id;
    2:i32 num;
}

struct ActicityInfo{
    1:i64 gid;
    2:string startTime;
    3:string endTime;
    4:list<PrizePair> prizes;
    5:i32 cost;
}


service BaseService{
    BaseResp Register(1:UserInfo req)
    LoginResp Login(1:UserInfo req)
    BaseResp Join(1:JoinReq req)
    BaseResp CreateGroup(1:GroupInfo req)
    BaseResp Sign(1:SignReq req)
    MonthSignResp SignMonth(1:MonthSignReq req)
    ChooseResp Choose(1:ChooseReq req)
    ChooseSubmitResp ChooseSubmit(1:ChooseSubmitReq req)
    BaseResp AddActivity(1:ActicityInfo req)
    BaseResp AddPrize(1:PrizeInfo req)
}

