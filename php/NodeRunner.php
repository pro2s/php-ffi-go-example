<?php

namespace FFITest;

class NodeRunner implements RunnerInterface
{
    private $total = 0;
    private $combinations = [];
    private $time = 0;
    private $cmd = '';

    public function __construct(string $path, Model $product)
    {
        $this->cmd = "node $path $product->id 2>&1";
    }

    public function run()
    {
        $out = null;
        $error = null;
        $result = \exec($this->cmd, $out, $error);

        if ($result && ($out[0] ?? false)) {
            $data = \json_decode($out[0], true);

            $this->combinations = $data['combinations'] ?? [];
            $this->total = $this->combinations['total'];
            $this->time = $data['time'] ?? 0;
        }
    }

    public function getTime()
    {
        return $this->time;
    }

    public function getName(): string
    {
        return 'NODE';
    }

    public function print()
    {
        echo $this->getName() , ':', PHP_EOL;
        echo 'Total: ', $this->total, PHP_EOL;
        echo \implode(',', $this->combinations[Data::TEST_HASH] ?? []), PHP_EOL;
        echo 'Time: ', $this->time, PHP_EOL;
        echo PHP_EOL;
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
