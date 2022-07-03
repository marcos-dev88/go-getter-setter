<?php

class TestPhpFile
{
	/**
	 *
	 * @var string
	 * 
	 */

	private   $myName = '';

	private    $myOther = '';

	private  $myTest   =    0;

	public function getMyName()
	{
		return $this->myName;
	}

	public function getMyOther()
	{
		return $this->myOther;
	}

	public function getMyTest()
	{
		return $this->myTest;
	}

	public function setMyName(string $myName)
	{
		$this->myName = $myName;
	}

	public function setMyOther(string $myOther)
	{
		$this->myOther = $myOther;
	}

	public function setMyTest(int $myTest)
	{
		$this->myTest = $myTest;
	}









}