<?php

namespace FFITest;

require 'vendor/autoload.php';

$service = new Combine(new Model(1));

$start = microtime(true);

$combinations = $service->getCombinationsArray(data());
echo 'Init: ', count($combinations), PHP_EOL;
$combinations = $service->filterLinked($combinations);
echo 'Filtered: ', count($combinations), PHP_EOL;
$combinations = $service->formatCombinations($combinations);
echo 'Format: ', count($combinations), PHP_EOL;
$end = microtime(true);
echo 'Time: ', $end - $start, PHP_EOL;
