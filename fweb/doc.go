// web框架.
//
// 接口说明支持以下标签.
/*
   @Constraints | [name] | [min] | [max] | [not] | [regexp]

   @Description

   @MapTo

   @Method [GET,POST等等]

   @Parameter | [name] | [type] | [required] | [default] | [description]

   @Remarks

   @Response

   @SortKey

   @Url
*/
// @Parameter和@Constraints是用' | '分隔, 如果要某个留空, 那么就应该'|  |',
// 中间要留够位置, 不然会出现非预期的情况.
//
// [required]和[default]必须二选一.
//
// @Url和@MapTo是必须的, 不然会报错.
package fweb

import (
//
)
