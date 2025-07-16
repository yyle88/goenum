package goenumgen

type NamingModeEnum string

const NamingMode = NamingModeEnum("Naming")

func (NamingModeEnum) Prefix() NamingModeEnum {
	return "Naming" + "-" + "Prefix"
}
func (NamingModeEnum) Suffix() NamingModeEnum {
	return "Naming" + "-" + "Suffix"
}
func (NamingModeEnum) Middle() NamingModeEnum {
	return "Naming" + "-" + "Middle"
}
func (NamingModeEnum) Single() NamingModeEnum {
	return "Naming" + "-" + "Single"
}
func (NamingModeEnum) Enums() []NamingModeEnum {
	return []NamingModeEnum{
		NamingMode.Prefix(),
		NamingMode.Suffix(),
		NamingMode.Middle(),
		NamingMode.Single(),
	}
}
