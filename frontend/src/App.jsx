import { useEffect, useState } from "react";
import { SparklesIcon, ArrowPathIcon } from "@heroicons/react/24/outline";

const Review = ({ score }) => {
  const stars = [];

  for (let i = 0; i < score; i++) {
    stars.push("⭐");
  }

  for (let i = score; i < 5; i++) {
    stars.push("☆");
  }

  return (
    <div>
      {stars.map((star, index) => (
        <span key={index}>{star}</span>
      ))}
    </div>
  );
};

function App() {
  const [isLoading, setIsLoading] = useState(false);
  const [reviews, setReviews] = useState([]);
  const [summarize, setSummarize] = useState(null);
  const [replies, setReplies] = useState([]);
  const [isLoadingReply, setIsLoadingReply] = useState([]);

  const handleSummarize = () => {
    setIsLoading(true);
    setSummarize(null);
    fetch("http://localhost:8081/query")
      .then((res) => res.json())
      .then((res) => {
        setSummarize(res);
        setIsLoading(false);
      })
      .catch(() => {
        setSummarize(null);
        setIsLoading(false);
      });
  };
  const handleReply = (index) => {
    const loadingReply = [...isLoadingReply];
    loadingReply[index] = true;
    setIsLoadingReply(loadingReply);

    setReplies((prev) => [...prev]);
    fetch(`http://localhost:8081/reply/${index}`)
      .then((res) => res.json())
      .then((res) => {
        loadingReply[index] = false;
        setIsLoadingReply(loadingReply);

        const newReply = [...replies];
        newReply[index] = res;
        setReplies(newReply);
      })
      .catch(() => {
        loadingReply[index] = false;
        setIsLoadingReply(loadingReply);

        setReplies((prev) => [...prev]);
      });
  };
  console.log(replies, "_debug");
  useEffect(() => {
    setIsLoading(true);
    fetch("http://localhost:8081/list")
      .then((res) => res.json())
      .then((res) => {
        setIsLoading(false);
        setReviews(res);
      })
      .catch(() => {
        setIsLoading(false);
        setReviews([]);
      });
  }, []);

  return (
    <section className="mx-auto container p-6 flex flex-col gap-12">
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 items-start">
        {reviews.map((review, index) => (
          <div
            key={index}
            className="bg-white border border-neutral-400 rounded shadow-md p-6 flex flex-col gap-6 justify-between"
          >
            <div className="flex flex-col gap-2">
              <span className="font-bold">{review?.author}</span>
              <Review score={review?.score} />
              <p>{review?.text}</p>
              {replies?.[index] && (
                <div className="italic text-sm">{replies?.[index]?.data}</div>
              )}
            </div>
            <div className="flex flex-row gap-3 items-center">
              {!replies?.[index] && (
                <button
                  className="w-fit flex flex-row gap-1 items-center cursor-pointer hover:text-fuchsia-400"
                  onClick={() => handleReply(index)}
                >
                  <SparklesIcon className="size-4" />
                  <span className="text-sm">Genera risposta</span>
                </button>
              )}
              {isLoadingReply[index] && (
                <ArrowPathIcon className="size-4 animate-spin" />
              )}
            </div>
          </div>
        ))}
      </div>

      <div>
        <div className="flex flex-row gap-3 items-center">
          {!summarize && (
            <button
              className="w-fit flex flex-row gap-1 items-center cursor-pointer hover:text-fuchsia-400 outline-0"
              onClick={handleSummarize}
            >
              <SparklesIcon className="size-4" />
              <span className="text-sm">Genera report recensioni</span>
            </button>
          )}
          {isLoading && <ArrowPathIcon className="size-4 animate-spin" />}
        </div>
        {summarize && <div className="font-semibold">{summarize?.data}</div>}
      </div>
    </section>
  );
}

export default App;
