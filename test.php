<?php

namespace FFITest;

require 'vendor/autoload.php';

$service = new GoCombine(new Model(1), __DIR__ . "/go/libutil.so");
$goRunner = new Runner($service);
$goRunner->run();
$goRunner->print();


$service = new Combine(new Model(1));
$phpRunner = new Runner($service);
$phpRunner->run();
$phpRunner->print();

$errors = [];
foreach ($phpRunner->getCombinations() as $combination) {
    if (!$goRunner->hasCombination($combination['hash'])) {
        $errors[] = $combination['hash'];
    }
}

echo count($errors) === 0 ? 'OK' : 'ERRORS: ' . implode(', ', $errors), PHP_EOL;
echo round($phpRunner->getTime() / $goRunner->getTime()), ' times', PHP_EOL;
