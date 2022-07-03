<?php

class MyPhhEight {

	private string $delorean;

	private int $clockTowerHour;

	private double $milesPerHour;

	private string $kaike;

	public function getDelorean()
	{
		return $this->delorean;
	}

	public function getClockTowerHour()
	{
		return $this->clockTowerHour;
	}

	public function getMilesPerHour()
	{
		return $this->milesPerHour;
	}

	public function setDelorean(string $delorean)
	{
		$this->delorean = $delorean;
	}

	public function setClockTowerHour(int $clockTowerHour)
	{
		$this->clockTowerHour = $clockTowerHour;
	}

	public function setMilesPerHour(double $milesPerHour)
	{
		$this->milesPerHour = $milesPerHour;
	}



	public function getKaike()
	{
		return $this->kaike;
	}

	public function setKaike(string $kaike)
	{
		$this->kaike = $kaike;
	}

}