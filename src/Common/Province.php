<?php

namespace Wulfheart\Alexander\Common;

class Province
{
    public string $payload;
    
    /**
     * Set the value of payload
     *
     * @return  self
     */ 
    public function setPayload($payload)
    {
        $this->payload = $payload;
    
        return $this;
    }
    
    /**
     * 
     * @return array [Province sup, Province sub]
     */
    public function split(): array {
        $split = explode("/", $this->payload);
        if(count($split) > 0){
            // $sup = 
        }
        if(count($split) > 1){

        }
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
