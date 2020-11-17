<?php

namespace Wulfheart\Alexander\Common;

class InconsistencyOrderTypeCount
{
    public OrderType $orderType;
    public int $found;
    public int $want;

    public function error(): string{
        return sprintf("InconsistencyOrderTypeCount:%s:Found:%d:Want:%d", $this->orderType->payload, $this->found, $this->want);
    }
}
