<?php

namespace Wulfheart\Alexander\Common;

class ErrDoubleBuild
{
    /**
     * 
     * @var array<Province>
     */
    public array $provinces;

    public function error(): string{
        // TODO: return fmt.Sprintf("ErrDoubleBuild:%v", self.Provinces)
        return sprintf("ErrDoubleBuild");
    }
}
