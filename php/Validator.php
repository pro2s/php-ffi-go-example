<?php

namespace FFITest;

class Validator
{
    private $main;
    private $check;

    public function __construct(RunnerInterface $main, array $check)
    {
        $this->main = $main;
        $this->check = $check;
    }

    public function validate()
    {
        $errors = \array_reduce(
            $this->check,
            fn (array $acc, RunnerInterface $runner) => $acc + [$runner->getName() => []],
            []
        );

        foreach ($this->main->getCombinations() as $combination) {
            foreach ($this->check as $runner) {
                if (!$runner->hasCombination($combination['hash'])) {
                    $errors[$runner->getName] = $combination['hash'];
                }
            }
        }

        foreach ($this->check as $runner) {
            $key = $runner->getName();
            echo $key, ': ', count($errors[$key]) === 0 ? 'OK' : 'ERRORS: ' . implode(', ', $errors[$key]), PHP_EOL;
            echo $runner->getName(), ': ', round($this->main->getTime() / $runner->getTime()), ' times', PHP_EOL;
        }
    }
}
