// Code generated by thriftgo (0.3.1). DO NOT EDIT.

package choose

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

type ChooseService interface {
	Choose(ctx context.Context, req *Empty) (r *Empty, err error)
}

type ChooseServiceClient struct {
	c thrift.TClient
}

func NewChooseServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ChooseServiceClient {
	return &ChooseServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewChooseServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ChooseServiceClient {
	return &ChooseServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewChooseServiceClient(c thrift.TClient) *ChooseServiceClient {
	return &ChooseServiceClient{
		c: c,
	}
}

func (p *ChooseServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *ChooseServiceClient) Choose(ctx context.Context, req *Empty) (r *Empty, err error) {
	var _args ChooseServiceChooseArgs
	_args.Req = req
	var _result ChooseServiceChooseResult
	if err = p.Client_().Call(ctx, "Choose", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type ChooseServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ChooseService
}

func (p *ChooseServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ChooseServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ChooseServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewChooseServiceProcessor(handler ChooseService) *ChooseServiceProcessor {
	self := &ChooseServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("Choose", &chooseServiceProcessorChoose{handler: handler})
	return self
}
func (p *ChooseServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type chooseServiceProcessorChoose struct {
	handler ChooseService
}

func (p *chooseServiceProcessorChoose) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ChooseServiceChooseArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Choose", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := ChooseServiceChooseResult{}
	var retval *Empty
	if retval, err2 = p.handler.Choose(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Choose: "+err2.Error())
		oprot.WriteMessageBegin("Choose", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Choose", thrift.REPLY, seqId); err2 != nil {
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

type ChooseServiceChooseArgs struct {
	Req *Empty `thrift:"req,1" frugal:"1,default,Empty" json:"req"`
}

func NewChooseServiceChooseArgs() *ChooseServiceChooseArgs {
	return &ChooseServiceChooseArgs{}
}

func (p *ChooseServiceChooseArgs) InitDefault() {
	*p = ChooseServiceChooseArgs{}
}

var ChooseServiceChooseArgs_Req_DEFAULT *Empty

func (p *ChooseServiceChooseArgs) GetReq() (v *Empty) {
	if !p.IsSetReq() {
		return ChooseServiceChooseArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *ChooseServiceChooseArgs) SetReq(val *Empty) {
	p.Req = val
}

var fieldIDToName_ChooseServiceChooseArgs = map[int16]string{
	1: "req",
}

func (p *ChooseServiceChooseArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChooseServiceChooseArgs) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ChooseServiceChooseArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ChooseServiceChooseArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewEmpty()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *ChooseServiceChooseArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Choose_args"); err != nil {
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

func (p *ChooseServiceChooseArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *ChooseServiceChooseArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ChooseServiceChooseArgs(%+v)", *p)
}

func (p *ChooseServiceChooseArgs) DeepEqual(ano *ChooseServiceChooseArgs) bool {
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

func (p *ChooseServiceChooseArgs) Field1DeepEqual(src *Empty) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type ChooseServiceChooseResult struct {
	Success *Empty `thrift:"success,0,optional" frugal:"0,optional,Empty" json:"success,omitempty"`
}

func NewChooseServiceChooseResult() *ChooseServiceChooseResult {
	return &ChooseServiceChooseResult{}
}

func (p *ChooseServiceChooseResult) InitDefault() {
	*p = ChooseServiceChooseResult{}
}

var ChooseServiceChooseResult_Success_DEFAULT *Empty

func (p *ChooseServiceChooseResult) GetSuccess() (v *Empty) {
	if !p.IsSetSuccess() {
		return ChooseServiceChooseResult_Success_DEFAULT
	}
	return p.Success
}
func (p *ChooseServiceChooseResult) SetSuccess(x interface{}) {
	p.Success = x.(*Empty)
}

var fieldIDToName_ChooseServiceChooseResult = map[int16]string{
	0: "success",
}

func (p *ChooseServiceChooseResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChooseServiceChooseResult) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ChooseServiceChooseResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ChooseServiceChooseResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewEmpty()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *ChooseServiceChooseResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Choose_result"); err != nil {
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

func (p *ChooseServiceChooseResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *ChooseServiceChooseResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ChooseServiceChooseResult(%+v)", *p)
}

func (p *ChooseServiceChooseResult) DeepEqual(ano *ChooseServiceChooseResult) bool {
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

func (p *ChooseServiceChooseResult) Field0DeepEqual(src *Empty) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
