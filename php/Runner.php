<?php

namespace FFITest;

class Runner
{
    protected $service;
    protected $start;
    protected $end;
    protected $combinations;

    public function __construct(CombineInterface $service)
    {
        $this->service = $service;
    }

    public function run()
    {
        $this->start = microtime(true);
        $data = $this->service->map(Data::getData());
        $this->combinations = $this->service->combine($data);
        $this->end = microtime(true);
    }

    public function getTime()
    {
        return $this->end - $this->start;
    }

    public function print()
    {
        echo $this->service->getName() . ':', PHP_EOL;
        echo 'Total: ', count($this->combinations), PHP_EOL;
        echo $this->service->formatCombination($this->combinations, Data::TEST_HASH), PHP_EOL;
        echo 'Time: ', $this->getTime(), PHP_EOL;
    }

    public function getCombinations(): array
    {
        return $this->combinations;
    }

    public function hasCombination(string $hash): bool
    {
        return isset($this->combinations[$hash]);
    }
}
