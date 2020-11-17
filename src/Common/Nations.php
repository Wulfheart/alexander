<?php

namespace Wulfheart\Alexander\Common;

class Nations
{
    /**
     * 
     * @var array<Nation>
     */
    public array $nations;

    public function equal(Nations $o): bool {
        // TODO: Test
        $diff = array_udiff($this->nations, $o->nations, function(Nation $a, $b){
            if($a->payload == $b->payload){
                return 0;
            } else {
                return 1;
            }
        });
        return count($diff == 0);
    }

    public function len(): int {
        return count($this->nations);
    }

    public function less(int $i, int $j): int{
        throw new \Exception("NotImplemented");
    }

    public function swap(int $i, int $j){
        throw new \Exception("NotImplemented");
    }
}
