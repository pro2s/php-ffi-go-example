<?php

namespace FFITest;

interface RunnerInterface
{
    public function run();
    public function print();
    public function getCombinations();
    public function hasCombination(string $hash): bool;
    public function getTime();
    public function getName(): string;
}
