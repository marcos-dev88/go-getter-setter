<?php

class Some
{
    /**
     *
     * @var string
     * 
     */

    private $my_name = '';

    private $my_other = '';

    private $my_test = 0;

    private $mydouble = 0.0;

    private $myBool = true;


    public function getMyName()
    {
        return $this->my_name;
    }

    public function setMyName(string $myName)
    {
        $this->my_name = $myName;
    }
}
