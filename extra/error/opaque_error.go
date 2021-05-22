// 黑盒error
/*
1.Error are just value 实现出了Error接口的类型都可以认为是Error
2.Don't just check errors, handle them gracefully 不要只是检查错误，请妥善处理
	接助pkg/errors
	通过Wrap方法，包装error并带上我们需要的错误信息
	通过Cause方法将里层的错误还原


*/

package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func f1() error {
	return errors.Wrapf(fmt.Errorf("real error"), "message f1")
}

func f2() error {
	return errors.Wrapf(f1(), "message f2")
}

func main() {
	if err := f2(); err != nil {
		fmt.Printf("%+v", err)
	}
}
