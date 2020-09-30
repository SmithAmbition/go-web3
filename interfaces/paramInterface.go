package interfaces
type IParamEncoder interface{
	EncodeParam(input ...interface{}) ([]interface{}, error)
}
type IParamDecoder interface{
	DecodeParam(input ...interface{}) ([]interface{}, error)
}