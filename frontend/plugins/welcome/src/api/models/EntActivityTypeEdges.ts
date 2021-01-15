/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntActivities,
    EntActivitiesFromJSON,
    EntActivitiesFromJSONTyped,
    EntActivitiesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntActivityTypeEdges
 */
export interface EntActivityTypeEdges {
    /**
     * Activities holds the value of the activities edge.
     * @type {Array<EntActivities>}
     * @memberof EntActivityTypeEdges
     */
    activities?: Array<EntActivities>;
}

export function EntActivityTypeEdgesFromJSON(json: any): EntActivityTypeEdges {
    return EntActivityTypeEdgesFromJSONTyped(json, false);
}

export function EntActivityTypeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntActivityTypeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'activities': !exists(json, 'activities') ? undefined : ((json['activities'] as Array<any>).map(EntActivitiesFromJSON)),
    };
}

export function EntActivityTypeEdgesToJSON(value?: EntActivityTypeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'activities': value.activities === undefined ? undefined : ((value.activities as Array<any>).map(EntActivitiesToJSON)),
    };
}


