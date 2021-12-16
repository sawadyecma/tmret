package a

func f0() {}

func f0err() error {
	return nil
}

func f0err2() (error, error) {
	return nil, nil
}

func f1() int {
	return 0
}

func f1err() (
	int,
	error,
) {
	return 0, nil
}

func f2() (int, int) {
	return 0, 0
}

func f2err() (
	int, int, error,
) {
	return 0, 0, nil
}

func f3() (int, int, int) {
	return 0, 0, 0
}

func f3err() (int, int, int, error) {
	return 0, 0, 0, nil
}

func h() (a int, err error) {
	return 0, err
}
