namespace go sign.sign

struct Empty{}

service SignService{
    Empty Sign(1:Empty req)
}