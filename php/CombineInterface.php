<?php

namespace FFITest;

interface CombineInterface
{
    public function getName(): string;
    public function map(array $data);
    public function combine($data);
    public function formatCombination(array $combinations, $hash): string;
}
