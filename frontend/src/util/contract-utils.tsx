import EventManagerAbi from '../abis/EventManager.json';
import EventAbi from '../abis/Event.json';
import dotenv from 'dotenv';

// Load environment variables from .env file
dotenv.config({path:"../.env"});

if (!process.env.EVENT_MANAGER_ADDRESS) {
  throw new Error("Missing EVENT_MANAGER_ADDRESS in .env file");
}

export const CONTRACTS = {
  EventManager: {
    address: process.env.EVENT_MANAGER_ADDRESS,
    abi: EventManagerAbi,
  },
  Event: {
    abi: EventAbi,
  },
} as const;

export type EventManagerABI = typeof EventManagerAbi;
export type EventABI = typeof EventAbi;
