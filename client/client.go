// This file is auto-generated, don't edit it. Thanks.
/**
* @return uuid
*/
package client

import (
  "github.com/google/uuid"
)


func Uuid () (_result string) {
  // V4 基于随机数
  u4 := uuid.New()
  return u4.String()
}

