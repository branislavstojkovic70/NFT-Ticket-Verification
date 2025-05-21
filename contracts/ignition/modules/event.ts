import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import { parseEther } from "viem";

const ONE_MICRO_ETHER: bigint = parseEther("0.000001"); // 1e12 wei = 0.000001 ETH

const EventManagerModule = buildModule("EventManagerModule", (m) => {
    const eventManager = m.contract("EventManager");

    const uuid = m.getParameter("uuid", "abc-123");
    const title = m.getParameter("title", "MyEvent");
    const startDate = m.getParameter(
        "startDate",
        Math.floor(Date.now() / 1000) + 3600
    );
    const endDate = m.getParameter(
        "endDate",
        Math.floor(Date.now() / 1000) + 7200
    );
    const numTickets = m.getParameter("numberOfTickets", 100);
    const pricePerTicket = m.getParameter("pricePerTicket", parseEther("0.01"));

    const tx = m.call(
        eventManager,
        "createEvent",
        [uuid, title, startDate, endDate, numTickets, pricePerTicket],
        {
            value: ONE_MICRO_ETHER,
        }
    );

    console.log(tx);
    return { eventManager };
});

export default EventManagerModule;
