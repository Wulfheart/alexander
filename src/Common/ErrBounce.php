<?php

namespace Wulfheart\Alexander\Common;

class ErrBounce
{
    public Province $province;

    public function error(): string{
        return sprintf("ErrBounce:%s", $this->province->payload);
    }
}
