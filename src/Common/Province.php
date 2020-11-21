<?php

namespace Wulfheart\Alexander\Common;

class Province
{
    public string $payload;

    public function __construct(string $s)
    {
        $this->payload = $s;
    }
    

    
    /**
     * 
     * @return array [Province sup, Province sub]
     */
    public function split(): array {
        $split = explode("/", $this->payload);
        if(count($split) > 0){
            $sup = new Province($split[0]);
        }
        if(count($split) > 1){
            $sub = new Province($split[1]);
        }
        return [$sup, $sub];
    }

    public function join(Province $n): Province {
        if($n->payload != ""){
            return new Province(sprintf("%s/%s", $this->payload, $n->payload));
        }
        return $this;
    }

    public function super(): Province {
        // ! This could be wrong
        return $this->split()[0];
    }

    public function sub(): Province {
        return $this->split()[1];
    }

    public function contains(Province $p) {
        throw new \Exception("NotImplemented");
    }

}
