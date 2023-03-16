import ProfChenMsg from "../../components/ProfChenMsg/ProfChenMsg";
import SubscriptionForm from "../../components/SubscriptionForm/SubscriptionForm";
import style from "./style.module.scss";

const Subscription = () => {
	return (
		<div className={style.Subscription}>
			<SubscriptionForm />
		</div>
	);
};

export default Subscription;
