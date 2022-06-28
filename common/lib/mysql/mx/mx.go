package mx

func GetName(alias, raw string, with *WithTrait) string {
	if alias == "" {
		alias = raw
	}
	if with.IsWithBackquote() {
		alias = "`" + alias + "`"
	}
	if !with.IsQuery() {
		with.Reset()
	}
	return alias
}
