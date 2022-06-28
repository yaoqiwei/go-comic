package table

import "fehu/common/lib/mysql/mx"

type join struct {
	table         mx.Container
	joinType      mx.JoinType
	joinCondition mx.ConditionMix
}

type joins []*join
