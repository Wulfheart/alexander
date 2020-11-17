<?php

namespace Wulfheart\Alexander\Common;

class InconsistencyMismatchedSupporter
{
    public Province $supportee;

    public function error(): string{
        return sprintf("InconsistencyMismatchedSupporter:%s", $this->supportee->payload);
    }
}
