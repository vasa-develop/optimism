# @eth-optiism/fee-estimation 

## Overview

This package is designed to provide an easy way to estimate gas on OP chains.

Fee estimation on OP-chains has both an l2 and l1 component. By default tools such as Viem, Wagmi, Ethers, and Web3.js do not support the l1 component. They will support this soon but in meantime, this library can help estimate fees for transactions, or act as a reference.
As these tools add support for gas estimation natively this README will be updated with framework specific instructions.

For more detailed information about gas fees on Optimism's Layer 2, you can visit their [official documentation](https://community.optimism.io/docs/developers/build/transaction-fees/#the-l2-execution-fee).

## GasPriceOracle contract

- The l2 contract that can estimate l1Fees is called [GasPriceOracle](../contracts-bedrock/contracts/l2/GasPriceOracle.sol) contract. This library provides utils for interacting with it at a high level.
- The GasPriceOracle is [deployed to Optimism](https://optimistic.etherscan.io/address/0x420000000000000000000000000000000000000F) and other OP chains at a predeployed address of `0x420000000000000000000000000000000000000F`

This library provides a higher level abstraction over the gasPriceOracle

## Installation

```bash
pnpm install @eth-optimism/fee-estimation
```

```bash
npm install @eth-optimism/fee-estimation
```

```bash
yarn add @eth-optimism/fee-estimation
```

## Usage

### Import the package

```ts
import { estimateFees } from '@eth-optimism/fee-estimation';
```

### Basic Usage

```ts
const fees = await estimateFees({
  client: {
    chainId: 10,
    rpcUrl: 'https://mainnet.optimism.io',
  },
  blockNumber: BigInt(106889079),
  account: '0xe371815c5f8a4f9acd1576879de288acd81269f1',
  to: '0xe35f24470730f5a488da9721548c1ab0b65b53d5',
  data: "0x5c19a95c00000000000000000000000046abfe1c972fca43766d6ad70e1c1df72f4bb4d1",
});
```

## API

### `estimateFees` function

```ts
estimateFees(params: EstimateFeeParams): Promise<bigint>
```

#### Parameters

`params`: An object with the following fields:

- `client`: An object with `rpcUrl` and `chainId` fields, or an instance of a Viem `PublicClient`.

- `blockNumber`: A BigInt representing the block number at which you want to estimate the fees.

- `blockTag`: A string representing the block tag to query from.

- `account`: A string representing the account making the transaction.

- `to`: A string representing the recipient of the transaction.

- `data`: A string representing the data being sent with the transaction. This should be a 0x-prefixed hex string.

#### Returns

A Promise that resolves to a BigInt representing the estimated fee in wei.

## FAQ

### How to encode function data?

You can use our package to encode the function data. Here is an example:

```ts
import { encodeFunctionData } from '@eth-optimism/fee-estimation';
import { optimistABI } from '@eth-optimism/contracts-ts';

const data = encodeFunctionData({
  functionName: 'burn',
  abi: optimistABI,
  args: [BigInt('0x77194aa25a06f932c10c0f25090f3046af2c85a6')],
});
```

This will return a 0x-prefixed hex string that represents the encoded function data.

## Testing

The package provides a set of tests that you can run to verify its operation. The tests are a great resource for examples. You can find them at `./src/estimateFees.spec.ts`.
