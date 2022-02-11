<?php

class TestPhpFile
{
    /**
     *
     * @var string
     * 
     */

    private $myAmazingVar = '';

    private $uhuuu = false;

    private $isWorking = true;
                                                                                                                                                                                                                                                                                                                                                                                                                                                                              
	public function getMyAmazingVar()
	{
		return $this->myAmazingVar;
	}

	public function isUhuuu()
	{
		return $this->uhuuu;
	}

	public function isIsWorking()
	{
		return $this->isWorking;
	}
   
	public function setMyAmazingVar(string $myAmazingVar)
	{
		$this->myAmazingVar = $myAmazingVar;
	}

	public function setUhuuu(bool $uhuuu)
	{
		$this->uhuuu = $uhuuu;
	}

	public function setIsWorking(bool $isWorking)
	{
		$this->isWorking = $isWorking;
	}

}