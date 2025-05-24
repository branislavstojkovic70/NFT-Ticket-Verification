import { ethers } from "ethers";
import { CONTRACTS } from "./contract-utils";

const RELAY_ENDPOINT = "http://localhost:8080/relay-tx";

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
  if (!(window as any).ethereum) throw new Error("MetaMask not available");

  try {
    const provider = new ethers.BrowserProvider((window as any).ethereum);
    const signer = await provider.getSigner();
    const address = await signer.getAddress();

    const gasPrice = await provider.getFeeData().then(data => data.gasPrice);
    if (!gasPrice) throw new Error("Could not fetch gas price");

    const contractInterface = new ethers.Interface(CONTRACTS.EventManager.abi);
    const data = contractInterface.encodeFunctionData("createEvent", [
      uuid,
      title,
      startDate,
      endDate,
      numberOfTickets,
      priceOfTicket
    ]);

    const estimatedGas = await provider.estimateGas({
      from: address,
      to: CONTRACTS.EventManager.address,
      data
    });
    
    const gasLimit = estimatedGas * 120n / 100n;

    const tx = {
      from: address,
      to: CONTRACTS.EventManager.address,
      data,
      value: ethers.parseUnits("0.000001", "ether"),
      gasLimit,
      gasPrice,
      chainId: (await provider.getNetwork()).chainId,
      nonce: await provider.getTransactionCount(address, "latest")
    };

    const signedTx = await signer.signTransaction(tx);
    
    const response = await fetch(RELAY_ENDPOINT, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ 
        signedTx,
        txData: { 
          method: "createEvent",
          params: { uuid, title, startDate, endDate, numberOfTickets, priceOfTicket }
        }
      })
    });

    if (!response.ok) {
      const error = await response.text();
      throw new Error(error);
    }

    const result = await response.json();
    return result.txHash;
  } catch (error) {
    console.error("Transaction preparation failed:", error);
    throw error;
  }
}