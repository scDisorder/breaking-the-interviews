package main

type Saver interface {
	Save(v any) error
}

func handleParallelTasks(s Saver, v []any) error {
	const workersLimit = 10

	errChan := make(chan error)
	defer close(errChan)

	semaphore := make(chan struct{}, workersLimit)
	defer close(semaphore)

	for _, payload := range v {
		go func(val any) {
			semaphore <- struct{}{}

			err1 := s.Save(val)

			<-semaphore

			errChan <- err1
		}(payload)
	}

	var errors []error
	for range v {
		if err := <-errChan; err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return errors[0]
	}

	return nil
}
