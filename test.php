<?php

namespace FFITest;

require 'vendor/autoload.php';

$service = new GoCombine(new Model(1), __DIR__ . "/go/libutil.so");

echo 'GO: ', PHP_EOL;
$start = microtime(true);
$goCombinations = $service->combine(data());
$end = microtime(true);
$go = $end - $start;
echo 'Time: ', $go, PHP_EOL;

$service = new Combine(new Model(1));

echo 'PHP: ', PHP_EOL;
$start = microtime(true);
$combinations = $service->getCombinationsArray(data());
echo 'Init: ', count($combinations), PHP_EOL;
$combinations = $service->filterLinked($combinations);
echo 'Filter: ', count($combinations), PHP_EOL;
$combinations = $service->formatCombinations($combinations);
echo 'Format: ', count($combinations), PHP_EOL;
$end = microtime(true);
$php = $end - $start;
echo 'Time: ', $php, PHP_EOL;

$errors = [];
foreach ($combinations as $combination) {
    if (!isset($goCombinations[$combination['hash']])) {
        $errors[] = $combination['hash'];
    }
}

echo !$errors ? 'OK' : implode(', ', $errors), ' ', round($php / $go), ' times',PHP_EOL;
