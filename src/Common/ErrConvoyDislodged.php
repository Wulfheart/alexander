<?php

namespace Wulfheart\Alexander\Common;

class ErrConvoyDislodged
{
    public Province $province;

    public function error(): string {
        return sprintf("ErrConvoyDislodged:%s", $this->province->payload);
    }
}
