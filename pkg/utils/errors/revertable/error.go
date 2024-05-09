package revertable

type RevertableError struct {
	baseErr error
	reason  string
}

var _ error = (*RevertableError)(nil)

func (r *RevertableError) Error() string {
	return "revertable error: " + r.baseErr.Error()
}

func (r *RevertableError) GetReason() string {
	return r.reason
}

func (r *RevertableError) Unwrap() error {
	return r.baseErr
}

func NewRevertable(err error, reason string) *RevertableError {
	return &RevertableError{
		baseErr: err,
		reason:  reason,
	}
}
