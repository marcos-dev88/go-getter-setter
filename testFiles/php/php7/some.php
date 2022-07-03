<?php

class Some
{
	/**
	 *
	 * @var string
	 * 
	 */

	private   $muName = '';

	private $myither = '';

	private $myuest = 0;


	public function getMuName()
	{
		return $this->muName;
	}

	public function getMyither()
	{
		return $this->myither;
	}

	public function getMyuest()
	{
		return $this->myuest;
	}

	public function setMuName(string $muName)
	{
		$this->muName = $muName;
	}

	public function setMyither(string $myither)
	{
		$this->myither = $myither;
	}

	public function setMyuest(int $myuest)
	{
		$this->myuest = $myuest;
	}









}