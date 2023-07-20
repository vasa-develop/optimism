// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const PreimageOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageLengths\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1001,\"contract\":\"contracts/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageParts\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1002,\"contract\":\"contracts/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimagePartOk\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_mapping(t_uint256,t_bytes32)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bytes32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var PreimageOracleStorageLayout = new(solc.StorageLayout)

var PreimageOracleDeployedBin = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c8063e15926111161005b578063e15926111461012a578063e52f0937146101a3578063fe4ac08e146101cf578063fef2b4ed146101fe5761007d565b806361238bde146100825780638542cf50146100b7578063e03110e1146100ee575b600080fd5b6100a56004803603604081101561009857600080fd5b508035906020013561021b565b60408051918252519081900360200190f35b6100da600480360360408110156100cd57600080fd5b5080359060200135610238565b604080519115158252519081900360200190f35b6101116004803603604081101561010457600080fd5b5080359060200135610258565b6040805192835260208301919091528051918290030190f35b6101a16004803603604081101561014057600080fd5b8135919081019060408101602082013564010000000081111561016257600080fd5b82018360208201111561017457600080fd5b8035906020019184600183028401116401000000008311171561019657600080fd5b50909250905061032b565b005b6100a5600480360360608110156101b957600080fd5b508035906020810135906040013560ff1661042b565b6101a1600480360360808110156101e557600080fd5b5080359060208101359060408101359060600135610525565b6100a56004803603602081101561021457600080fd5b503561058c565b600160209081526000928352604080842090915290825290205481565b600260209081526000928352604080842090915290825290205460ff1681565b6000828152600260209081526040808320848452909152812054819060ff166102e257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7072652d696d616765206d757374206578697374000000000000000000000000604482015290519081900360640190fd5b506000838152602081815260409091205460088101848301106103085783816008010391505b506000938452600160209081526040808620948652939052919092205492909150565b6044356000806008830186111561034157600080fd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b60006104368461059e565b60008181526002602090815260408083208380528252808320805460017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090911681179091558484528252808320838052909152908190209085901c60c085901b1790819055909150601860ff84161115610508576000828152600260209081526040808320828452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558584528252808320828452909152902060c085901b90555b50600081815260208190526040902060ff90921690915592915050565b6000838152600260209081526040808320878452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558684528252808320968352958152858220939093559283529082905291902055565b60006020819052908152604090205481565b7f01000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8216176105eb816105f1565b92915050565b600090815233602052604090207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f0100000000000000000000000000000000000000000000000000000000000000179056fea164736f6c6343000706000a"

var PreimageOracleDeployedSourceMap = "198:5963:16:-:0;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;410:68;;;;;;;;;;;;;;;;-1:-1:-1;410:68:16;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;553:66;;;;;;;;;;;;;;;;-1:-1:-1;553:66:16;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;890:564;;;;;;;;;;;;;;;;-1:-1:-1;890:564:16;;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;4697:1462;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;4697:1462:16;;-1:-1:-1;4697:1462:16;-1:-1:-1;4697:1462:16;:::i;:::-;;3319:1087;;;;;;;;;;;;;;;;-1:-1:-1;3319:1087:16;;;;;;;;;;;;;;:::i;1801:262::-;;;;;;;;;;;;;;;;-1:-1:-1;1801:262:16;;;;;;;;;;;;;;;;;:::i;292:50::-;;;;;;;;;;;;;;;;-1:-1:-1;292:50:16;;:::i;410:68::-;;;;;;;;;;;;;;;;;;;;;;;;:::o;553:66::-;;;;;;;;;;;;;;;;;;;;;;;;;;:::o;890:564::-;990:12;1043:20;;;:14;:20;;;;;;;;:29;;;;;;;;;990:12;;1043:29;;1035:62;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;1228:14:16;1245:21;;;1216:2;1245:21;;;;;;;;1305:1;1296:10;;1280:12;;;:26;1276:87;;1345:7;1332:6;1341:1;1332:10;:20;1322:30;;1276:87;-1:-1:-1;1419:19:16;;;;:13;:19;;;;;;;;:28;;;;;;;;;;;;890:564;;-1:-1:-1;890:564:16:o;4697:1462::-;4993:4;4980:18;4798:12;;5122:1;5112:12;;5096:29;;5093:2;;;5154:1;5151;5144:12;5093:2;5413:3;5409:14;;;5313:4;5397:27;5444:11;5418:4;5563:16;5444:11;5545:41;5776:29;;;5780:11;5776:29;5770:36;5828:20;;;;5975:19;5968:27;5997:11;5965:44;6028:19;;;;6006:1;6028:19;;;;;;;;:32;;;;;;;;:39;;;;6063:4;6028:39;;;;;;6077:18;;;;;;;;:31;;;;;;;;;:38;;;;6125:20;;;;;;;;;;;:27;;;;-1:-1:-1;;;;4697:1462:16:o;3319:1087::-;3434:12;3535:36;3564:6;3535:28;:36::i;:::-;3674:13;3970:20;;;:14;:20;;;;3892:2;3970:20;;;:23;;;;;;;;:30;;3996:4;3970:30;;;;;;;;;4010:19;;;;;;;;:22;;;;;;;;;;3888:14;;;;3875:3;3871:15;;;3868:35;4010:30;;;;3970:20;;-1:-1:-1;4146:2:16;3970:30;4138:10;;;4134:159;;;4164:13;4206:20;;;:14;:20;;;;;;;;:24;;;;;;;;:31;;;;4233:4;4206:31;;;;;;4251:19;;;;;;;;:23;;;;;;;;4189:3;4180:12;;;4251:31;;4134:159;-1:-1:-1;4370:15:16;:21;;;;;;;;;;:29;;;;;;;4386:4;3319:1087;-1:-1:-1;;3319:1087:16:o;1801:262::-;1934:19;;;;:14;:19;;;;;;;;:31;;;;;;;;:38;;;;1968:4;1934:38;;;;;;1982:18;;;;;;;;:30;;;;;;;;;:37;;;;2029:20;;;;;;;;;;:27;1801:262::o;292:50::-;;;;;;;;;;;;;;:::o;492:353:15:-;752:11;777:19;765:32;;749:49;824:14;749:49;824:8;:14::i;:::-;817:21;492:353;-1:-1:-1;;492:353:15:o;1222:430::-;1277:21;1426:15;;;1467:8;1461:4;1454:22;1595:4;1582:18;;1602:19;1578:44;1624:11;1575:61;;1319:327::o"

func init() {
	if err := json.Unmarshal([]byte(PreimageOracleStorageLayoutJSON), PreimageOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["PreimageOracle"] = PreimageOracleStorageLayout
	deployedBytecodes["PreimageOracle"] = PreimageOracleDeployedBin
}
