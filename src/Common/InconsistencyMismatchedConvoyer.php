<?php

namespace Wulfheart\Alexander\Common;

class InconsistencyMismatchedConvoyer
{
    public Province $convoyee;

    public function error(): string{
        return sprintf("InconsistencyMismatchedSupporter:%s", $this->convoyee->payload);
    }
}
