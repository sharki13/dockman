package wrapper

type MultipassWrapper struct{}

func MakeMultipassWrapper() Wrapper {
	return &MultipassWrapper{}
}

func (w *MultipassWrapper) GetType() WarpperType {
	return MultipassWrapperType
}

func (w *MultipassWrapper) GetInstances() ([]Instance, error) {
	return []Instance{}, nil
}
