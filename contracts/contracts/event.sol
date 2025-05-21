// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

///The sender is not the owner
error EventManager_NotOwner();

///Not enough ether provided: `amount`. Minimum required: `required`
error EventManager_InsufficientFunds(uint256 amount, uint256 required);

error EventManager_TransactionFailed();

/// Nothing to withdraw.
error EventManager_ZeroBalance();

///Not enough ether provided: `amount`. Ticket price is: `required`
error Event_InsufficientFunds(uint256 amount, uint256 required);

error Event_TransactionFailed();

/// All tickets for this event sold.
error Event_SoldOut();

/// The ticket has been used.
error Event_TicketUsed();

/// Event has started.
error Event_EventStarted();

/// Not your NFT.
error Event_NotOwner();

/// Nothing to withdraw.
error Event_ZeroBalance();

enum EventType { Conference, Workshop, Webinar }

struct EventData {
    string UUID;
    string Title;
    uint startDate;
    uint endDate;
    uint numberOfTickets;
    uint256 priceOfTicket;
}

contract EventManager {
    //Constants
    uint256 public constant PRICE = 1e12;
    
    //Public
    address public immutable owner;

    //EVENTS
    event Created(address _creator, address _event);

    //MAPPING
    mapping(address => EventData[]) public userEvents;

    constructor() {
        owner = msg.sender;
    }

    function createEvent(
        string memory _UUID,
        string memory _Title,
        uint _startDate,
        uint _endDate,
        uint _numberOfTickets,
        uint256 _priceOfTicket
    ) public payable {
        if (msg.value < PRICE){
            revert EventManager_InsufficientFunds({
                amount: msg.value,
                required: PRICE
            });
        }

        (bool success, ) = payable(address(this)).call{value: msg.value}("");

        if (!success) revert EventManager_TransactionFailed();

        EventData memory newEvent = EventData({
            UUID: _UUID,
            Title: _Title,
            startDate: _startDate,
            endDate: _endDate,
            numberOfTickets: _numberOfTickets,
            priceOfTicket: _priceOfTicket
        });
        Event newEventContract = new Event(msg.sender, newEvent);

        userEvents[msg.sender].push(newEvent);
        emit Created(msg.sender, address(newEventContract));
    }

    function withdraw () external onlyOwner {
        uint balance = address(this).balance;
        if (balance == 0) revert EventManager_ZeroBalance();
        (bool success, ) = payable(msg.sender).call{value: balance}("");
        if (!success) revert EventManager_TransactionFailed();
    }

    modifier onlyOwner() {
        if (msg.sender != owner){
            revert EventManager_NotOwner();
        }
        _;
    }
}

contract Event is ERC721 {
    address immutable public owner;
    EventData public eventData;
    uint public ticketCounter;

    mapping(uint => bool) ticketsUsed;

    event TicketPurchased(address indexed _buyer, uint ticket);

    constructor(
        address _owner,
        EventData memory _eventData
    )
        ERC721(_eventData.Title, getSymbol(_eventData.Title))
    {
        owner = _owner;
        eventData = _eventData;
    }

    function buyTicket() external payable eventNotExpired  {
        if (eventData.numberOfTickets == ticketCounter) revert Event_SoldOut();

        if (msg.value < eventData.priceOfTicket) 
            revert Event_InsufficientFunds({ amount: msg.value, required: eventData.priceOfTicket});

        (bool success, ) = payable(address(this)).call{value: msg.value}("");

        if (!success) revert Event_TransactionFailed();

        ticketCounter += 1;

        _safeMint(msg.sender, ticketCounter);
        ticketsUsed[ticketCounter] = false;

        emit TicketPurchased(msg.sender, ticketCounter);
    }

    function refundTicket(uint ticketId) external eventNotExpired {
        if (ticketsUsed[ticketId]) revert Event_TicketUsed();
        
        if (block.timestamp > eventData.startDate) revert Event_EventStarted();

        address ticketOwner = ownerOf(ticketId);
        if (ticketOwner != msg.sender) revert Event_NotOwner();

        delete ticketsUsed[ticketId];
        _burn(ticketId);
        (bool success, ) = payable(msg.sender).call{value: eventData.priceOfTicket}("");

        if (!success) revert Event_TransactionFailed();
    }

    function markAsUsed(uint ticketId) external eventNotExpired {
        if (ticketsUsed[ticketId]) revert Event_TicketUsed();
        ticketsUsed[ticketId] = true;
    }

    function withdraw () external eventExpired {
        uint balance = address(this).balance;
        if (balance == 0) revert Event_ZeroBalance();
        (bool success, ) = payable(msg.sender).call{value: balance}("");
        if (!success) revert Event_TransactionFailed();
    }

    function getSymbol(string memory _title) internal pure returns (string memory) {
        bytes memory b = bytes(_title);
        bytes memory result = new bytes(3);
        for (uint i = 0; i < 3 && i < b.length; i++) {
            result[i] = b[i];
        }
        return string(result);
    }

    modifier eventNotExpired() {
        require(block.timestamp < eventData.endDate, "Event has ended.");
        _;
    }

    modifier eventExpired() {
        require(block.timestamp > eventData.endDate, "Event has not ended.");
        _;
    }
}
