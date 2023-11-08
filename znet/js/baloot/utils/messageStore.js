const messageStore = (
  lastJsonMessage,
  setMessageHistory,
  setDeck,
  setPlayingCards,
  setGametake,
  setPlayerID,
  setPlayerTakes,
  setCards,
  setTakes,
  setScore
) => {
  console.log(lastJsonMessage);
  if (lastJsonMessage !== null) {
    if (lastJsonMessage !== {}) {
      switch (lastJsonMessage.id) {
        case 1:
          setPlayerID((prev) => lastJsonMessage.player.id);
          setCards((prev) => lastJsonMessage.player.hand.cards);
          setTakes((prev) => lastJsonMessage.availableTakes);
          break;
        case 2:
          setCards((prev) => []);
          setTakes((prev) => []);
          setPlayerTakes((prev) => []);
          setPlayingCards((prev) => lastJsonMessage.player);
          setGametake((prev) => lastJsonMessage.gametake);
          break;
        case 5:
          setPlayerTakes((prev) => [...prev, lastJsonMessage.take]);
          setTakes((prev) => lastJsonMessage.availableTakes);
          break;
        case 6:
          setDeck((prev) => lastJsonMessage.deck);
          setPlayingCards((prev) => lastJsonMessage.player);
          setScore((prev) => [lastJsonMessage.scoreTeamA, lastJsonMessage.scoreTeamB])
          break;
        default:
          throw new Error("Error message id not found");
      }
    }
    setMessageHistory((prev) => prev.concat(lastJsonMessage));
  }
};

export default messageStore;
