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
} from './estimateFees'
import { optimistABI, optimistAddress } from '@eth-optimism/contracts-ts'
import { parseGwei } from 'viem'

// using this optimist https://optimistic.etherscan.io/tx/0xaa291efba7ea40b0742e5ff84a1e7831a2eb6c2fc35001fa03ba80fd3b609dc9
const blockNumber = BigInt(107028270)
const optimistOwnerAddress = "0x77194aa25a06f932c10c0f25090f3046af2c85a6" as const
const functionCallData = {
  functionName: 'burn',
  // this is an erc721 abi
  abi: optimistABI,
  args: [BigInt(optimistOwnerAddress)],
  account: optimistOwnerAddress,
  to: optimistAddress[10],
  chainId: 10
} as const

const callDataWithPriorityFees = {
  ...functionCallData,
  maxFeePerGas:parseGwei('2'),
  maxPriorityFeePerGas: parseGwei('2'),
}

const clientParams = {
  chainId: functionCallData.chainId,
  rpcUrl: process.env.VITE_L2_RPC_URL ?? 'https://mainnet.optimism.io',
} as const

const viemClient = getL2Client(clientParams)

const paramsWithRpcUrl = {
  client: clientParams,
  blockNumber,
} as const

const paramsWithViemClient = {
  client: viemClient,
  viemClient,
  blockNumber,
} as const

beforeEach(() => {
  vi.resetAllMocks()
})

test('estimateFees should return correct fees', async () => {
  const res = await estimateFees({ ...paramsWithRpcUrl, ...functionCallData })
  expect(res).toMatchInlineSnapshot('20573185261089n')
  expect(await estimateFees({ ...paramsWithRpcUrl, ...functionCallData })).toMatchInlineSnapshot('20573185261089n')
  expect(await estimateFees({ ...paramsWithRpcUrl, ...functionCallData })).toMatchInlineSnapshot('20573185261089n')
  expect(await estimateFees({ ...paramsWithRpcUrl, ...functionCallData })).toMatchInlineSnapshot('20573185261089n')
  expect(formatEther(res)).toMatchInlineSnapshot('"0.000020573185261089"')
  // what is the l2 and l1 part of the fees for reference?
  const l1Fee = await getL1Fee({ ...paramsWithRpcUrl, ...functionCallData })
  const l2Fee = res - l1Fee
  expect(l1Fee).toMatchInlineSnapshot('20573185216764n')
  expect(formatEther(l1Fee)).toMatchInlineSnapshot('"0.000020573185216764"')
  expect(l2Fee).toMatchInlineSnapshot('44325n')
  expect(formatEther(l2Fee)).toMatchInlineSnapshot('"0.000000000000044325"')
})

test('baseFee should return the correct result', async () => {
  expect(await baseFee(paramsWithRpcUrl)).toMatchInlineSnapshot('64n')
  expect(await baseFee(paramsWithViemClient)).toMatchInlineSnapshot('64n')
})

test('decimals should return the correct result', async () => {
  expect(await decimals(paramsWithRpcUrl)).toMatchInlineSnapshot('6n')
  expect(await decimals(paramsWithViemClient)).toMatchInlineSnapshot('6n')
})

test('gasPrice should return the correct result', async () => {
  expect(await gasPrice(paramsWithRpcUrl)).toMatchInlineSnapshot('64n')
  expect(await gasPrice(paramsWithViemClient)).toMatchInlineSnapshot('64n')
})

test('getL1Fee should return the correct result', async () => {
  expect(await getL1Fee({ ...paramsWithRpcUrl, ...functionCallData })).toMatchInlineSnapshot('20573185216764n')
  expect(await getL1Fee({ ...paramsWithViemClient, ...functionCallData })).toMatchInlineSnapshot('20573185216764n')
  expect(await getL1Fee({ ...paramsWithViemClient, ...callDataWithPriorityFees })).toMatchInlineSnapshot('21536974073765n')
  expect(formatEther(await getL1Fee({ ...paramsWithViemClient, ...functionCallData }))).toMatchInlineSnapshot('"0.000020573185216764"')
})

test('getL1GasUsed should return the correct result', async () => {
  expect(await getL1GasUsed({ ...paramsWithRpcUrl, ...functionCallData })).toMatchInlineSnapshot('2220n')
  expect(await getL1GasUsed({ ...paramsWithViemClient, ...functionCallData })).toMatchInlineSnapshot('2220n')
  expect(await getL1GasUsed({ ...paramsWithViemClient, ...callDataWithPriorityFees })).toMatchInlineSnapshot('2324n')
})

test('l1BaseFee should return the correct result', async () => {
  expect(await l1BaseFee(paramsWithRpcUrl)).toMatchInlineSnapshot('13548538813n')
  expect(await l1BaseFee(paramsWithViemClient)).toMatchInlineSnapshot('13548538813n')
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

