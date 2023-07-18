/**
 * The first 2 test cases are good documentation of how to use this library
 */
import { vi, test, expect, beforeEach } from 'vitest'
import { formatEther } from 'viem/utils'
import {
  baseFee,
  decimals,
  estimateFees,
  gasPrice,
  getL1Fee,
  getL1GasUsed,
  getL2Client,
  l1BaseFee,
  overhead,
  scalar,
  version,
  encodeFunctionData
} from './estimateFees'
import { optimistABI, optimistAddress } from '@eth-optimism/contracts-ts'

// https://optimistic.etherscan.io/tx/0xceb34708f1de2ea2336dbdd81ca2cb7b12ec2ab6adc86a0d7d81c8c4416f81bb
const data =
  '0x5c19a95c00000000000000000000000046abfe1c972fca43766d6ad70e1c1df72f4bb4d1'
const blockNumber = BigInt(106889079)


const clientParams = {
  chainId: 10,
  rpcUrl: process.env.VITE_L2_RPC_URL ?? 'https://mainnet.optimism.io',
} as const

const viemClient = getL2Client(clientParams)

const paramsWithRpcUrl = {
  client: clientParams,
  blockNumber
}

const paramsWithViemClient = {
  client: viemClient,
  viemClient,
  blockNumber
}

beforeEach(() => {
  vi.resetAllMocks()
})

test('estimateFees should return correct fees', async () => {
  const estimateFeesParams = {
    data:
      '0xd1e16f0a603acf1f8150e020434b096e408bafa429a7134fbdad2ae82a9b2b882bfcf5fe174162cf4b3d5f2ab46ff6433792fc99885d55ce0972d982583cc1e11b64b1d8d50121c0497642000000000000000000000000000000000000060a2c8052ed420000000000000000000000000000000000004234002c8052edba12222222228d8ba445958a75a0704d566bf2c84200000000000000000000000000000000000006420000000000000000000000000000000000004239965c9dab5448482cf7e002f583c812ceb53046000100000000000000000003',
    account: '0xe371815c5f8a4f9acd1576879de288acd81269f1',
    to: '0xe35f24470730f5a488da9721548c1ab0b65b53d5',
  } as const
  const res = await estimateFees({ ...paramsWithRpcUrl, ...estimateFeesParams })
  expect(res).toMatchInlineSnapshot('33237088573005n')
  expect(await estimateFees({ ...paramsWithViemClient, ...estimateFeesParams })).toMatchInlineSnapshot('33237088573005n')
  expect(formatEther(res)).toMatchInlineSnapshot('"0.000033237088573005"')
})

test('users should be able to generate data using encodeFunctionData', async () => {
  // using this optimist https://optimistic.etherscan.io/tx/0xaa291efba7ea40b0742e5ff84a1e7831a2eb6c2fc35001fa03ba80fd3b609dc9
  const blockNumber = BigInt(107028270)
  const optimistOwnerAddress = "0x77194aa25a06f932c10c0f25090f3046af2c85a6" as const
  const client = getL2Client(clientParams)
  const data = encodeFunctionData({
    // not sure why you would estimate gas of ownerOf but this is just an example
    functionName: 'burn',
    // this is an erc721 abi
    abi: optimistABI,
    args: [BigInt(optimistOwnerAddress)],
  })
  expect(data).toMatchInlineSnapshot('"0x42966c6800000000000000000000000077194aa25a06f932c10c0f25090f3046af2c85a6"')
  const res = await estimateFees({
    client,
    data,
    account: optimistOwnerAddress,
    to: optimistAddress[10],
    blockNumber
  })
  expect(res).toMatchInlineSnapshot('15828378580466n')
})

test('baseFee should return the correct result', async () => {
  expect(await baseFee(paramsWithRpcUrl)).toMatchInlineSnapshot('65n')
  expect(await baseFee(paramsWithViemClient)).toMatchInlineSnapshot('65n')
})

test('decimals should return the correct result', async () => {
  expect(await decimals(paramsWithRpcUrl)).toMatchInlineSnapshot('6n')
  expect(await decimals(paramsWithViemClient)).toMatchInlineSnapshot('6n')
})

test('gasPrice should return the correct result', async () => {
  expect(await gasPrice(paramsWithRpcUrl)).toMatchInlineSnapshot('65n')
  expect(await gasPrice(paramsWithViemClient)).toMatchInlineSnapshot('65n')
})

test('getL1Fee should return the correct result', async () => {
  expect(await getL1Fee(data, paramsWithRpcUrl)).toMatchInlineSnapshot('15130316373984n')
  expect(await getL1Fee(data, paramsWithViemClient)).toMatchInlineSnapshot('15130316373984n')
})

test('getL1GasUsed should return the correct result', async () => {
  expect(await getL1GasUsed(data, paramsWithRpcUrl)).toMatchInlineSnapshot('1708n')
  expect(await getL1GasUsed(data, paramsWithViemClient)).toMatchInlineSnapshot('1708n')
})

test('l1BaseFee should return the correct result', async () => {
  expect(await l1BaseFee(paramsWithRpcUrl)).toMatchInlineSnapshot('12951022000n')
  expect(await l1BaseFee(paramsWithViemClient)).toMatchInlineSnapshot('12951022000n')
})

test('overhead should return the correct result', async () => {
  expect(await overhead(paramsWithRpcUrl)).toMatchInlineSnapshot('188n')
  expect(await overhead(paramsWithViemClient)).toMatchInlineSnapshot('188n')
})

test('scalar should return the correct result', async () => {
  expect(await scalar(paramsWithRpcUrl)).toMatchInlineSnapshot('684000n')
  expect(await scalar(paramsWithViemClient)).toMatchInlineSnapshot('684000n')
})

test('version should return the correct result', async () => {
  expect(await version(paramsWithRpcUrl)).toMatchInlineSnapshot('"1.0.0"')
  expect(await version(paramsWithViemClient)).toMatchInlineSnapshot('"1.0.0"')
})

