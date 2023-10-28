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
) => {
  if (lastJsonMessage !== null) {
    if (lastJsonMessage !== {}) {
      switch (lastJsonMessage.id) {
        case 1:
          setPlayerID((prev) => lastJsonMessage.player.id);
          setCards((prev) => lastJsonMessage.player.hand.Cards);
          setTakes((prev) => lastJsonMessage.available_takes);
          break;
        case 2:
          setCards((prev) => []);
          setTakes((prev) => []);
          setPlayerTakes((prev) => []);
          setPlayingCards((prev) => lastJsonMessage.player.playing_hand.Cards);
          setGametake((prev) => lastJsonMessage.gametake.Name);
          break;
        case 5:
          setPlayerTakes((prev) => [...prev, lastJsonMessage.take]);
          setTakes((prev) => lastJsonMessage.available_takes);
          break;
        case 6:
          setDeck((prev) => lastJsonMessage.deck);
          setPlayingCards((prev) => lastJsonMessage.player.playing_hand.Cards);
          break;
        default:
          throw new Error("Error message id not found");
      }
    }
    setMessageHistory((prev) => prev.concat(lastJsonMessage));
  }
};

export default messageStore;
