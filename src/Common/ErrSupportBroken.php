<?php

namespace Wulfheart\Alexander\Common;

class ErrSupportBroken
{
    public Province $province;

    public function error(): string {
        return sprintf("ErrSupportBroken:%s", $this->province->payload);
    }
}
