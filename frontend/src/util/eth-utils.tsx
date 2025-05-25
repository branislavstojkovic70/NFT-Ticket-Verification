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

    const contract = new ethers.Contract(CONTRACTS.EventManager.address, CONTRACTS.EventManager.abi, signer);
    
    // Populate transaction
    const txRequest = await contract.createEvent.populateTransaction(
      uuid,
      title,
      startDate,
      endDate,
      numberOfTickets,
      priceOfTicket
    );

    txRequest.from = address;
    txRequest.value = ethers.parseUnits("0.000001", "ether");
    txRequest.nonce = await provider.getTransactionCount(address, "latest");
    txRequest.chainId = (await provider.getNetwork()).chainId;
    txRequest.gasLimit = await provider.estimateGas(txRequest);
    txRequest.gasPrice = (await provider.getFeeData()).gasPrice!;

    const unsignedTx = ethers.Transaction.from(txRequest);
    const serializedUnsignedTx = unsignedTx.unsignedSerialized;

    const signature = await signer.signMessage(ethers.getBytes(serializedUnsignedTx));

    const response = await fetch(RELAY_ENDPOINT, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        serializedUnsignedTx,
        signature,
        txData: {
          method: "createEvent",
          params: { uuid, title, startDate, endDate, numberOfTickets, priceOfTicket },
        },
      }),
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
