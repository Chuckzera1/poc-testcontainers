package user

func (d *deleteUserUseCase) Delete(id uint64) error {
	return d.repository.Delete(id)
}
