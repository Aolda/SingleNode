// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract AoldaClient {
  event CallAolda(string fn, string[] arguments);
  mapping(bytes32 => string) private result;

  function callAolda(string memory fn, string[] memory arguments) public {
    emit CallAolda(fn, arguments);
  }

  function setValue(bytes32 signature, string memory value) public{
      result[signature]=value;
  }

  function getValue(bytes32 signature) public view returns(string memory){
      return result[signature];
  }

  function makeSignature(string memory fn, string[] memory arguments) public pure returns(bytes32) {
    string memory concated = fn;
    for(uint256 i=0;i<arguments.length;i++){
      concated= string(abi.encodePacked(concated,arguments[i]));
    }
    return stringToBytes32(concated);
  }

    function stringToBytes32(string memory source) internal pure returns (bytes32 result) {
        // require(bytes(source).length <= 32); // causes error
        // but string have to be max 32 chars
        // https://ethereum.stackexchange.com/questions/9603/understanding-mload-assembly-function
        // http://solidity.readthedocs.io/en/latest/assembly.html
        assembly {
        result := mload(add(source, 32))
        }
    }//
}