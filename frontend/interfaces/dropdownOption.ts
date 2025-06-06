import { ReactNode } from "react";
import PropTypes from "prop-types";

export default PropTypes.shape({
  disabled: PropTypes.bool,
  label: PropTypes.string,
  value: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
});

export type TooltipContent = ReactNode;

export interface IDropdownOption {
  disabled?: boolean;
  label: string | JSX.Element;
  value: string | number;
  helpText?: ReactNode;
  premiumOnly?: boolean;
  tooltipContent?: TooltipContent;
  onClick?: (() => void) | void | null;
}
