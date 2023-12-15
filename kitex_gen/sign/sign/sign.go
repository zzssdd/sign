// Code generated by thriftgo (0.3.1). DO NOT EDIT.

package sign

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type Empty struct {
}

func NewEmpty() *Empty {
	return &Empty{}
}

func (p *Empty) InitDefault() {
	*p = Empty{}
}

var fieldIDToName_Empty = map[int16]string{}

func (p *Empty) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Empty) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("Empty"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Empty) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Empty(%+v)", *p)
}

func (p *Empty) DeepEqual(ano *Empty) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	return true
}

type SignService interface {
	Sign(ctx context.Context, req *Empty) (r *Empty, err error)
}

type SignServiceClient struct {
	c thrift.TClient
}

func NewSignServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SignServiceClient {
	return &SignServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewSignServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SignServiceClient {
	return &SignServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewSignServiceClient(c thrift.TClient) *SignServiceClient {
	return &SignServiceClient{
		c: c,
	}
}

func (p *SignServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *SignServiceClient) Sign(ctx context.Context, req *Empty) (r *Empty, err error) {
	var _args SignServiceSignArgs
	_args.Req = req
	var _result SignServiceSignResult
	if err = p.Client_().Call(ctx, "Sign", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type SignServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      SignService
}

func (p *SignServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *SignServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *SignServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewSignServiceProcessor(handler SignService) *SignServiceProcessor {
	self := &SignServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("Sign", &signServiceProcessorSign{handler: handler})
	return self
}
func (p *SignServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type signServiceProcessorSign struct {
	handler SignService
}

func (p *signServiceProcessorSign) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SignServiceSignArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Sign", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := SignServiceSignResult{}
	var retval *Empty
	if retval, err2 = p.handler.Sign(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Sign: "+err2.Error())
		oprot.WriteMessageBegin("Sign", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Sign", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type SignServiceSignArgs struct {
	Req *Empty `thrift:"req,1" frugal:"1,default,Empty" json:"req"`
}

func NewSignServiceSignArgs() *SignServiceSignArgs {
	return &SignServiceSignArgs{}
}

func (p *SignServiceSignArgs) InitDefault() {
	*p = SignServiceSignArgs{}
}

var SignServiceSignArgs_Req_DEFAULT *Empty

func (p *SignServiceSignArgs) GetReq() (v *Empty) {
	if !p.IsSetReq() {
		return SignServiceSignArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *SignServiceSignArgs) SetReq(val *Empty) {
	p.Req = val
}

var fieldIDToName_SignServiceSignArgs = map[int16]string{
	1: "req",
}

func (p *SignServiceSignArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SignServiceSignArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SignServiceSignArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SignServiceSignArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewEmpty()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *SignServiceSignArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Sign_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SignServiceSignArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SignServiceSignArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SignServiceSignArgs(%+v)", *p)
}

func (p *SignServiceSignArgs) DeepEqual(ano *SignServiceSignArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *SignServiceSignArgs) Field1DeepEqual(src *Empty) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type SignServiceSignResult struct {
	Success *Empty `thrift:"success,0,optional" frugal:"0,optional,Empty" json:"success,omitempty"`
}

func NewSignServiceSignResult() *SignServiceSignResult {
	return &SignServiceSignResult{}
}

func (p *SignServiceSignResult) InitDefault() {
	*p = SignServiceSignResult{}
}

var SignServiceSignResult_Success_DEFAULT *Empty

func (p *SignServiceSignResult) GetSuccess() (v *Empty) {
	if !p.IsSetSuccess() {
		return SignServiceSignResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SignServiceSignResult) SetSuccess(x interface{}) {
	p.Success = x.(*Empty)
}

var fieldIDToName_SignServiceSignResult = map[int16]string{
	0: "success",
}

func (p *SignServiceSignResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SignServiceSignResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SignServiceSignResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SignServiceSignResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewEmpty()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *SignServiceSignResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Sign_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SignServiceSignResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *SignServiceSignResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SignServiceSignResult(%+v)", *p)
}

func (p *SignServiceSignResult) DeepEqual(ano *SignServiceSignResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *SignServiceSignResult) Field0DeepEqual(src *Empty) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
