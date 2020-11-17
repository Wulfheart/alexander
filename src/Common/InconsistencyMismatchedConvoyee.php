<?php

namespace Wulfheart\Alexander\Common;

class InconsistencyMismatchedConvoyee
{
    public Province $convoyer;

    public function error(): string{
        return sprintf("InconsistencyMismatchedSupporter:%s", $this->convoyer->payload);
    }
}
