<?php

namespace Wulfheart\Alexander\Common;

class Province
{
    public string $payload;

    /**
     * 
     * @return array [Province sup, Province sub]
     */
    public function split(): array {
        throw new \Exception("NotImplemented");
    }

    public function join(Province $n): Province {
        throw new \Exception("NotImplemented");
    }

    public function super(): Province {
        throw new \Exception("NotImplemented");
    }

    public function sub(): Province {
        throw new \Exception("NotImplemented");
    }

    public function contains(Province $p) {
        throw new \Exception("NotImplemented");
    }
}
