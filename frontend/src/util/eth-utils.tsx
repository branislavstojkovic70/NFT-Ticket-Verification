import { ethers } from "ethers";

const EVENT_MANAGER_ADDRESS = "0xYourContractAddress";

// export async function estimateGasLimit({
//     providerUrl,
//     contractAddress,
//     abi,
//     methodName,
//     methodArgs,
//     fromAddress
//   }: {
//     providerUrl: string;
//     contractAddress: string;
//     abi: any[];
//     methodName: string;
//     methodArgs: any[];
//     fromAddress: string;
//   }) {
//     try {
//       const provider = new ethers.JsonRpcProvider(providerUrl);
//       const contract = new ethers.Contract(contractAddress, abi, provider);
//       const estimatedGas = await contract.est[methodName](...methodArgs, {
//         from: fromAddress
//       });
  
//       // Optionally add buffer
//       const gasLimitWithBuffer = estimatedGas.mul(ethers.toBigInt(1.1)); // +10%
//       return gasLimitWithBuffer;
//     } catch (error) {
//       console.error('Gas estimation failed:', error);
//       throw error;
//     }
//   }

export async function prepareSignedCreateEventTx({
  uuid,
  title,
  startDate,
  endDate,
  numberOfTickets,
  priceOfTicket
}: {
  uuid: string;
  title: string;
  startDate: number;
  endDate: number;
  numberOfTickets: number;
  priceOfTicket: bigint;
}) {
  if (!window.ethereum) throw new Error("MetaMask not available");

  const provider = new ethers.BrowserProvider(window.ethereum);
  const signer = await provider.getSigner();
  const address = await signer.getAddress();

  const contractInterface = new ethers.Interface(EventManagerAbi);
  const data = contractInterface.encodeFunctionData("createEvent", [
    uuid,
    title,
    startDate,
    endDate,
    numberOfTickets,
    priceOfTicket
  ]);

  const gasPrice = 1000;

  const tx = {
    from: address,
    to: EVENT_MANAGER_ADDRESS,
    data,
    value: ethers.parseUnits("0.000001", "ether"), // PRICE
    0,
    gasLimit: 300000n,
    gasPrice
  };

  // Potpisivanje
  const rawTx = await signer.signTransaction(tx);

  // Šalji na server
  const res = await fetch("http://localhost:8080/relay-tx", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ rawTx })
  });

  const text = await res.text();
  if (!res.ok) throw new Error(text);

  return text; // tx hash, npr.
}
