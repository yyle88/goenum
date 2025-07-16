package enums

type StatusEnum int

const Status = StatusEnum(0)

func (StatusEnum) OK() StatusEnum {
	return 200
}
func (StatusEnum) WA() StatusEnum {
	return 400
}
func (StatusEnum) Enums() []StatusEnum {
	return []StatusEnum{
		Status.OK(),
		Status.WA(),
	}
}
